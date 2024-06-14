<script>
export default {
    data: function () {
        return {

        }
    },
    props: ["commentId", "photoAuthorId", "userId", "photoId", "comment", "isPhotoOwner"],
    methods: {
        async deleteComment() {
            try {
                let response = await this.$axios.delete(`/users/${this.$props.photoAuthorId}/photos/${this.$props.photoId}/comments/${this.$props.commentId}`, {
                    headers: { "Authorization": `Bearer ${localStorage.getItem("token")}` },
                });
                this.$emit('commentDeleted', this.$props.commentId)
            } catch (e) {
                this.errormsg = e.toString();
            }
        }
    }
}
</script>

<template>
    <div>
        <h5>{{ this.$props.userId }} says:</h5>
        <p>{{ this.$props.comment }}</p>
        <button v-if="this.$props.isPhotoOwner" @click="deleteComment">Delete Comment</button>
    </div>
</template>
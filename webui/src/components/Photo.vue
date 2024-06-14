<script>
export default {
    data: function () {
        return {
            isLikedByBearer: false,
            reactiveComments: [],
            reactiveLikes: [],
            commentText: "",
            activeUser: "",
            photoURL: "",
            showComments: false,
        }
    },
    props: ["authorId", "photoId", "date", "likes", "comments", "isAuthor"],
    methods: {
        getActiveUser() {
            this.activeUser = localStorage.getItem("token") || "";
        },
        loadInfo() {
            this.reactiveLikes = this.$props.likes;
            this.reactiveComments = this.$props.comments;
            this.isLikedByBearer = this.$props.likes.some(x => x.user_id == this.activeUser)
            this.photoURL = __API_URL__ + "/users/" + this.$props.authorId + "/photos/" + this.$props.photoId
        },
        async postComment(authorId, photoId) {
            try {
                let response = await this.$axios.post(`/users/${authorId}/photos/${photoId}/comments`, {
                    "comment": this.commentText
                }, {
                    headers: { "Authorization": `Bearer ${this.activeUser}` },
                });
                this.reactiveComments.push(response.data)
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.commentText = ""
        },
        async addLike() {
            try {
                let response = await this.$axios.put(`/users/${this.$props.authorId}/photos/${this.$props.photoId}/likes/${this.activeUser}`, {}, {
                    headers: { "Authorization": `Bearer ${this.activeUser}` },
                });
                this.reactiveLikes.push({ "user_id": this.activeUser })
                this.isLikedByBearer = true;
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async removeLike() {
            try {
                let response = await this.$axios.delete(`/users/${this.$props.authorId}/photos/${this.$props.photoId}/likes/${this.activeUser}`, {
                    headers: { "Authorization": `Bearer ${this.activeUser}` },
                });
                this.isLikedByBearer = false;
                this.reactiveLikes = this.reactiveLikes.filter(x => x.user_id != this.activeUser)
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async deletePhoto() {
            try {
                let response = await this.$axios.delete(`/users/${this.$props.authorId}/photos/${this.$props.photoId}`, {
                    headers: { "Authorization": `Bearer ${this.activeUser}` },
                });
                this.isLikedByBearer = false;
                this.reactiveLikes = this.reactiveLikes.filter(x => x.user_id != this.activeUser)
                this.$emit('photoDeleted', this.$props.photoId)
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        deleteComment(commentId) {
            this.reactiveComments = this.reactiveComments.filter(x => x.comment_id != commentId)
        }
    },
    mounted() {
        this.getActiveUser();
        this.loadInfo();
    }
}
</script>

<template>
    <div class="card mb-4" style="width: fit-content; margin-inline: auto;">
        <div class="card-body">
            <img :src="photoURL" style="width: 32rem; height: 18rem; object-fit: contain;"/>
            <h5 class="card-title">{{ this.$props.authorId }}</h5>
            <div>{{ this.reactiveLikes.length }} Likes, {{ this.reactiveComments.length }} Comments</div>
            <div>
                <button type="button" class="button" @click="addLike" v-if="!isLikedByBearer">Like</button>
                <button type="button" class="button" @click="removeLike" v-else>Remove Like</button>
                <button type="button" class="button" @click="deletePhoto" v-if="this.$props.isAuthor">Delete Photo</button>
                <button type="button" class="button" @click="() => showComments = !showComments">Comment</button>
            </div>
            <div v-if="showComments">
                <PhotoComment
                    v-for="comment in reactiveComments"
                    :key="comment.comment_id"
                    :commentId="comment.comment_id"
                    :photoAuthorId="this.$props.authorId"
                    :userId="this.activeUser"
                    :photoId="this.$props.photoId"
                    :comment="comment.comment"
                    :isPhotoOwner="this.activeUser == this.$props.authorId"
                    @commentDeleted="deleteComment"
                />
                <form @submit.prevent="postComment(this.$props.authorId, this.$props.photoId)">
                    <input type="text" name="commentText" v-model="commentText" placeholder="Write a comment..." />
                    <button type="submit">Submit Comment</button>
                </form>
            </div>
        </div>
    </div>
</template>
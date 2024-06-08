<script>
export default {
    data: function () {
        return {
            errormsg: null,
            following: 0,
            followers: 0,
            activeUser: "",
            profile: {
                user_id: "",
                followers: [],
                followings: [],
                photos: []
            },
            followersCounter: 0,
            followingsCounter: 0,
            isFollower: false,
            isBanned: false,
        }
    },
    methods: {
        async fetchUser() {
            this.errormsg = null;
            try {
                let response = await this.$axios.get(`/users/${this.$route.params.userId}`, {
                    headers: { "Authorization": `Bearer ${localStorage.getItem("token")}` }
                });
                this.profile = response.data;
                this.isFollower = this.profile.followers.some((follow) => follow.user_id == this.activeUser)
                this.followersCounter = response.data.followers.length;
                this.followingsCounter = response.data.followings.length;
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async followUser() {
            this.errormsg = null;
            try {
                let response = await this.$axios.put(`/users/${this.$route.params.userId}/followers/${this.activeUser}`, {}, {
                    headers: { "Authorization": `Bearer ${localStorage.getItem("token")}` }
                });
                this.followersCounter += 1;
                this.isFollower = true;
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async unfollowUser() {
            this.errormsg = null;
            try {
                let response = await this.$axios.delete(`/users/${this.$route.params.userId}/followers/${this.activeUser}`, {
                    headers: {
                        "Authorization": `Bearer ${localStorage.getItem("token")}`
                    }
                });
                this.followersCounter -= 1;
                this.isFollower = false;
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async banUser() {
            this.errormsg = null;
            try {
                let response = await this.$axios.put(`/users/${this.activeUser}/banned/${this.$route.params.userId}`, {}, {
                    headers: {
                        "Authorization": `Bearer ${localStorage.getItem("token")}`
                    }
                });
                this.isBanned = true;
                if(this.isFollower) {
                    this.followersCounter -= 1;
                    this.isFollower = false;
                }
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async unbanUser() {
            this.errormsg = null;
            try {
                let response = await this.$axios.delete(`/users/${this.activeUser}/banned/${this.$route.params.userId}`, {
                    headers: {
                        "Authorization": `Bearer ${localStorage.getItem("token")}`
                    }
                });
                this.isBanned = false;
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        getActiveUser() {
            this.activeUser = localStorage.getItem("token")
        },
    },
    async mounted() {
        this.getActiveUser();
        await this.fetchUser();
    }
}
</script>

<template>
    <div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div>
            <h1>Profile View</h1>
            <h3>@{{ profile.user_id || this.$route.params.userId }}</h3>
            <p>Followers {{ followersCounter }}</p>
            <p>Followings {{ followingsCounter }}</p>
        </div>
        <div v-if="activeUser !== $route.params.userId">
            <button v-if="!isFollower && !isBanned" type="button" @click="followUser">Follow</button>
            <button v-if="isFollower" type="button" @click="unfollowUser">UnFollow</button>
            <button v-if="!isBanned" type="button" @click="banUser">Ban</button>
            <button v-if="isBanned" type="button" @click="unbanUser">UnBan</button>
        </div>
        <div v-else>
            <form enctype="multipart/form-data">
                <input type="file" name="file" />
            </form>
        </div>
    </div>
</template>
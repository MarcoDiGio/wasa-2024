<script>
export default {
    data: function () {
        return {
            errormsg: null,
            username: "",
            activeUser: "",
        }
    },
    methods: {
        async changeUsername() {
            this.errormsg = null;
            try {
                let response = await this.$axios.put(`/users/${this.activeUser}`, {
                    // it still uses replace, in case space.prevent doesn't work
                    // or in case someone pastes text with spaces for username
                    user_id: this.username.replace(/\s/g, ''),
                }, {
                    headers: { "Authorization": `Bearer ${this.activeUser}` }
                });
                localStorage.setItem("token", response.data.user_id)
                this.$router.replace("/")
            } catch (e) {
                console.log(e)
                this.errormsg = e.toString();
            }
            this.username = "";
        },
    },
    mounted() {
        this.activeUser = localStorage.getItem("token")
        if (this.activeUser == "") {
            this.$axios.replace('/')
        }
    }
}
</script>

<template>
    <div>
        <h1>SETTINGS VIEW</h1>
        <form @submit.prevent="changeUsername">
            <input type="text" v-model="username" @keydown.space.prevent required minlength="3" maxlength="16" placeholder="New username here..."/>
            <button type="submit">Invia</button>
            <div><b>Insert your new username here</b></div>
            <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        </form>
    </div>
</template>
<script>
export default {
    data: function () {
        return {
            errormsg: null,
            loading: false,
            username: "",
        }
    },
    methods: {
        async doLogin() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session", {
                    // it still uses replace, in case space.prevent doesn't work
                    // or in case someone pastes text with spaces for username
                    user_id: this.username.replace(/\s/g, ''),
                });
                localStorage.setItem("token", response.data.user_id)
                this.$router.replace("/")
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
            this.username = "";
        }
    }
}
</script>

<template>
    <div>
        <h1>LOGIN VIEW</h1>
        <form @submit.prevent="doLogin">
            <input type="text" v-model="username" @keydown.space.prevent required minlength="3" maxlength="16" />
            <button type="submit">Invia</button>
            <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        </form>
    </div>
</template>
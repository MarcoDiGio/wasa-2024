<script>
export default {
	data: function () {
		return {
			errormsg: "",
			posts: [],
			users: [],
			activeUser: "",
			searchResults: [],
			userToSearch: "",
			commentText: "",
		}
	},
	methods: {
		getActiveUser() {
			this.activeUser = localStorage.getItem("token") || "";
		},
		async searchUser() {
			this.errormsg = null;
			try {
				let response = await this.$axios.get(`/users/${this.userToSearch}/search`, {
					headers: { "Authorization": `Bearer ${this.activeUser}` },
				});
				this.searchResults = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async getStream() {
			try {
				let response = await this.$axios.get(`/users/${this.activeUser}/stream`, {
					headers: { "Authorization": `Bearer ${this.activeUser}` },
				});
				this.posts = response.data;
				console.log(this.posts)
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async postComment(authorId, photoId) {
			try {
				let response = await this.$axios.post(`/users/${authorId}/photos/${photoId}/comments`, {
					"comment": this.commentText
				}, {
					headers: { "Authorization": `Bearer ${this.activeUser}` },
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		goToProfile() {
			this.$router.replace('/users/' + this.activeUser)
		}
	},
	async mounted() {
		this.getActiveUser();
		if (this.activeUser !== "") {
			await this.getStream()
		} else {
			this.$router.replace('/login')
		}
	},

}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="goToProfile">
						Profile
					</button>
				</div>
			</div>
		</div>
		<form @submit.prevent="searchUser">
			<input type="text" pattern="[a-zA-Z0-9]{3,16}" v-model="userToSearch" placeholder="Search an user..." />
			<button type="submit">Search User</button>
		</form>
		<h3>Hello, {{ activeUser }}</h3>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
	<section v-if="searchResults.length > 0" class="d-flex flex-wrap flex-md-nowrap align-items-center mt-4 mb-4">
		<div v-for="user in searchResults" :key="user.user_id" class="card search-result" @click="this.$router.replace(`/users/${user.user_id}`)">
			<div class="card-body">
				<h5 class="card-title">{{ user.user_id }}</h5>
			</div>
		</div>
	</section>
	<section>
		<Photo 
			v-for="(photo, index) in posts"
			:key="index"
			:photoId="photo.photo_ID"
			:authorId="photo.author_id"
			:date="photo.date"
			:likes="photo.likes"
			:comments="photo.comments"
			:isAuthor="photo.author_id == activeUser"
		/>
	</section>
</template>

<style>
.search-result {
	cursor: pointer;
	transition: opacity 0.5s ease-out;
}

.card+.card {
	margin-left: 2rem;
}

.search-result:hover {
	opacity: 0.5;
}

.button+.button {
	margin-left: .5rem;
}

</style>

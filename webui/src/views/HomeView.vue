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
	<section class="d-flex flex-wrap flex-md-nowrap align-items-center mr-2">
		<div v-for="user in searchResults" class="card" @click="this.$router.replace(`/users/${user.user_id}`)">
			<div class="card-body">
				<h5 class="card-title">{{ user.user_id }}</h5>
			</div>
		</div>
		<!-- @TODO: CONVERT THIS DIV INTO PHOTO COMPONENT -->
		<div v-for="post in posts" class="card">
			<div class="card-body">
				<img :src="'http://localhost:3000/users/' + post.author_id + '/photos/' + post.photo_ID" class="img" />
				<h5 class="card-title">{{ post.author_id }}</h5>
				<div>{{ post.likes.length }} Likes, {{ post.comments.length }} Comments</div>
				<div>
					<button type="button" class="button">Like</button>
					<button type="button" class="button">Comment</button>
				</div>
				<div v-for="comment in post.comments">
					<h5>{{ comment.user_id }} says:</h5>
					<p>{{ comment.content }}</p>
				</div>
				<form @submit.prevent="postComment(post.author_id, post.photo_ID)">
					<input type="text" name="commentText" v-model="commentText" placeholder="Write a comment..."/>
					<button type="submit">Submit Comment</button>
				</form>
			</div>
		</div>
	</section>
</template>

<style>
.card {
	cursor: pointer;
	transition: opacity 0.5s ease-out;
}

.card+.card {
	margin-left: 2rem;
}

.card:hover {
	opacity: 0.5;
}

.button+.button {
	margin-left: .5rem;
}

.img {
	max-width: 15rem;
	object-fit: cover;
}
</style>

import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SettingsView from '../views/SettingsView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/login', component: LoginView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
		{path: '/users/:userId', component: ProfileView},
		{path: '/users/:userId/settings', component: SettingsView}
	]
})

export default router

import { createRouter, createWebHistory } from "vue-router";
import Login from "./components/pages/Login.vue";
import Signup from "./components/pages/Signup.vue";
// import Dashboard from "./components/pages/Dashboard.vue";

// pages
import Dashboard from "./components/pages/Dashboard.vue";
import Task from "./components/pages/Task.vue";

const routes = [
	{
		path: "/login",
		component: Login,
		name: "Login",
	},
	{
		path: "/signup",
		component: Signup,
		name: "Signup",
	},
	{
		path: "/",
		component: Dashboard,
		name: "Dashboard",
		// meta: {
		// 	requiresAuth: true,
		// },
	},
	{
		path: "/task/:id",
		component: Task,
		name: "Task",
		// meta: {
		// 	requiresAuth: true,
		// },
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

router.beforeEach((to, from, next) => {
	const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

	if (requiresAuth) {
		const jwtToken = localStorage.getItem("jwtToken");

		if (!jwtToken) {
			next({ name: "Login" });
		} else {
			next();
		}
	} else {
		next();
	}
});

export default router;

import { createRouter, createWebHistory } from "vue-router";

// pages
// pages goes here

const routes = [
	{
		path: "/login",
		component: Login,
		name: "Login",
	},
	{
		path: "/",
		component: Dashboard,
		name: "Dashboard",
		meta: {
			requiresAuth: true,
		},
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

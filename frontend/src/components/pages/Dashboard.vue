<script setup>
	import Navbar from "../utils/Navbar.vue";
	import { ref, onMounted } from "vue";
	import axios from "axios";
	import "primeicons/primeicons.css";
	import NavbarTop from "../utils/NavbarTop.vue";
	import AddProjectForm from "../utils/AddProjectForm.vue";
	const showAddTaskForm = ref(false);
	const showAddTaskFormModal = () => {
		showAddTaskForm.value = true;
	};
	const closeFormModal = () => {
		showAddTaskForm.value = false;
	};
	const isCollapsed = ref(false);
	const toggleCollapse = () => {
		isCollapsed.value = !isCollapsed.value;
	};

	const projects = ref([]);

	const dataProject = async () => {
		const token = localStorage.getItem("token");
		const user_id = localStorage.getItem("userID");

		try {
			const res = await axios.get(
				`http://localhost:8080/onize/projects/user/${user_id}`
			);

			projects.value = res.data;
		} catch (err) {
			console.log(err);
		}
	};

	onMounted(dataProject);

	const formatDate = (dateString) => {
		const options = { day: "numeric", month: "numeric", year: "numeric" };
		const date = new Date(dateString);
		return date.toLocaleDateString("en-GB", options);
	};
</script>

<template>
	<div v-if="showAddTaskForm" class="overlay" @click="closeFormModal"></div>
	<AddProjectForm
		v-if="showAddTaskForm"
		@close-form-modal="closeFormModal"
		:class="{ 'fade-scale-in': showAddTaskForm }"
	/>

	<div class="container">
		<Navbar :isCollapsed="isCollapsed" @toggleCollapse="toggleCollapse" />
		<div
			class="dashboard-content"
			:style="{ marginLeft: isCollapsed ? '80px' : '247px' }"
		>
			<NavbarTop />
			<!-- navbar-top -->
			<div class="nav-top">
				<div class="left-nav">Project List</div>
				<div class="right-nav">
					<i class="pi pi-plus" @click="showAddTaskFormModal"></i>
				</div>
			</div>

			<router-link
				:to="{ name: 'Task', params: { id: project.id } }"
				class="task-container"
				v-for="project in projects"
				:key="project.id"
				style="text-decoration: none; color: inherit"
			>
				<div class="task">
					<div class="top-task">
						<div class="task-judul">
							<h3>{{ project.title }}</h3>
						</div>
						<div class="task-date">
							<i class="pi pi-calendar"></i>
							{{ formatDate(project.created_at) }}
						</div>
						<div class="task-action">
							<i class="pi pi-ellipsis-v"></i>
						</div>
					</div>
					<div class="task-badge">
						<i class="pi pi-circle-fill green"></i>
						<p>Study</p>
					</div>
				</div>
			</router-link>
		</div>
	</div>
</template>

<style scoped>
	hr {
		width: 100%;
		color: #222;
	}

	.task-date {
		display: flex;
		margin-right: 10%;
		width: 17%;
		display: inline-block;
		/* background-color: grey; */
		border: 1px solid #222;
		border-radius: 100px;
		padding: 7px;
		text-align: center;
	}

	.red {
		color: red;
	}

	.yellow {
		color: yellow;
	}

	.green {
		color: green;
	}

	.task-action {
		cursor: pointer;
	}

	.task-container {
		margin-top: 15px;
		display: flex;
		width: 98%;
		justify-content: center;
	}

	.task-checklist {
		margin: 0;
		align-items: center;
		display: flex;
		justify-content: center;
		margin-right: 20px;
	}

	.task-badge {
		display: flex;
		gap: 11px;
		width: 100%;
		margin-left: 80px;
		align-items: center;
		padding: 0;
	}

	.task-badge i {
		font-size: 12px;
		margin: 0;
		margin-top: 5px;
	}

	.task-badge p {
		margin: 0;
		margin-top: 5px;
		font-size: 12px;
	}

	.task-judul {
		width: 100%;
	}

	.task-judul h3 {
		margin: 0;
		margin-bottom: 5px;
		font-weight: 400;
	}

	.task {
		display: flex;
		width: 100%;
		flex-direction: column;
		padding: 20px;
		transition: all 0.3s;
		border-radius: 10px;
		/* margin-top: 10px; */
		align-items: center;
		justify-content: center;
	}

	.task:hover {
		background: #6427aa;
	}

	.top-task {
		display: flex;
		justify-content: space-between;
		width: 100%;
		align-items: center;
		height: 20px;
	}

	.nav-top {
		display: flex;
		justify-content: space-between;
		width: 100%;
		padding: 1.7%;
	}

	.left-nav {
		display: flex;
		align-items: center;
		gap: 10px;
	}

	.right-nav {
		color: #6427aa;
		padding-right: 10px;
	}

	.right-nav i {
		cursor: pointer;
		font-size: 20px;
		transition: transform 0.3s ease;
	}

	.right-nav i:hover {
		transform: rotate(90deg);
	}

	.container {
		display: flex;
		width: 100vw;
	}

	.collapsed-container {
		transition: margin-left 0.5s ease-in-out;
	}

	.dashboard-content {
		width: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		flex: 1;
		transition: margin-left 0.5s ease-in-out;
	}

	.card-container {
		width: 97%;
		display: flex;
		justify-content: center;
		flex-direction: column;
		align-items: center;
	}

	.header {
		align-self: flex-start;
	}

	.card {
		width: 100%;
		height: 60%;
		overflow: hidden;
		background-color: #6427aa;
	}

	.stats-container {
		display: flex;
		width: 100%;
		justify-content: space-evenly;
		gap: 20px;
	}

	.stats {
		text-align: start;
	}

	.stats h3 {
		color: white;
		margin: 0;
		font-weight: 600;
	}

	.stats p {
		font-size: 4rem;
		margin: 0;
		font-weight: 700;
		color: white;
	}

	.fade-scale-in {
		animation: fadeScaleIn 0.4s;
	}

	@keyframes fadeScaleIn {
		from {
			opacity: 0;
			transform: scale(0.8);
		}

		to {
			opacity: 1;
			transform: scale(1);
		}
	}

	.overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.5);
		z-index: 100;
	}
</style>

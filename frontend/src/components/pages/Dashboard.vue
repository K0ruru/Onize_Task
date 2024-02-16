<script setup>
	import Navbar from "../utils/Navbar.vue";
	import { ref } from "vue";

	import Card from "primevue/card";

	const isCollapsed = ref(false);

	const toggleCollapse = () => {
		isCollapsed.value = !isCollapsed.value;
	};
</script>

<template>
	<div class="container" :class="{ 'collapsed-container': isCollapsed }">
		<Navbar :isCollapsed="isCollapsed" @toggleCollapse="toggleCollapse" />
		<div class="dashboard-content">
			<div class="card-container">
				<div class="header">
					<h1>Welcome to Onize!</h1>
				</div>
				<Card class="card">
					<template #content>
						<div class="tasks-container">
							<div class="tasks">
								<h3>Total Task</h3>
								<p>120</p>
							</div>
							<div class="tasks">
								<h3>Total Task</h3>
								<p>120</p>
							</div>
							<div class="tasks">
								<h3>Total Task</h3>
								<p>120</p>
							</div>
							<div class="tasks">
								<h3>Total Task</h3>
								<p>120</p>
							</div>
						</div>
					</template>
				</Card>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.container {
		display: flex;
		width: 1920px;
	}

	.collapsed-container {
		transition: margin-left 0.5s ease-in-out;
	}

	.dashboard-content {
		width: 100vw;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.card-container {
		width: 80%;
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
		overflow: hidden;
		background-color: rgb(164, 9, 211);
	}

	.tasks-container {
		display: flex;
		width: 100%;
		justify-content: space-evenly;
	}

	.tasks {
		text-align: start;
	}

	.tasks h3 {
		color: white;
		margin: 0;
		font-weight: 600;
	}

	.tasks p {
		font-size: 5rem;
		margin: 0;
		font-weight: 700;
		color: white;
	}
</style>

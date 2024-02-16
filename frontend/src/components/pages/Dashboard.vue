<script setup>
	import Navbar from "../utils/Navbar.vue";
	import { ref } from "vue";

	const isCollapsed = ref(false);

	const toggleCollapse = () => {
		isCollapsed.value = !isCollapsed.value;
	};
</script>

<template>
	<div class="container" :class="{ 'collapsed-container': isCollapsed }">
		<Navbar :isCollapsed="isCollapsed" @toggleCollapse="toggleCollapse" />
	</div>
</template>

<style scoped>
	.container {
		display: flex;
	}

	.collapsed-container {
		transition: margin-left 0.5s ease;
	}
</style>

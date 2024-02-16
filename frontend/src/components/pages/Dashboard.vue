<script setup>
import Navbar from "../utils/Navbar.vue";
import { ref } from "vue";
import Card from "primevue/card";
import Dropdown from 'primevue/dropdown';
import Message from 'primevue/message';

const isCollapsed = ref(false);
const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value;
};
</script>

<template>
  <div class="container" :class="{ 'collapsed-container': isCollapsed }">
    <Navbar :isCollapsed="isCollapsed" @toggleCollapse="toggleCollapse" />
    <div class="dashboard-content">

      <div class="nav-top">

        <div class="left-nav">
          <p>Jan 2024</p>
          <Dropdown v-model="selectedCity" :options="cities" optionLabel="name" placeholder="Week" class="w-full md:w-14rem" />
        </div>
        <div class="right-nav">
          <p>| |</p>
        </div>

      </div>

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
.nav-top {
  display: flex;
  justify-content: space-between;
  width: 100%;
  border-bottom: 1px solid #222;
  padding:1.7%;
}

.left-nav {
  display: flex;
  align-items: center;
  gap: 20px;
}

.right-nav {
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

.tasks-container {
  display: flex;
  width: 100%;
  justify-content: space-evenly;
  gap: 20px;
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
  font-size: 4rem;
  margin: 0;
  font-weight: 700;
  color: white;
}
</style>

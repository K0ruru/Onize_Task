<script setup>
import "primeicons/primeicons.css";
// import { defineProps, defineEmits } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const { isCollapsed } = defineProps(["isCollapsed"]);
const emit = defineEmits();

const handleCollapse = () => {
  emit("toggleCollapse");
};

const logout = () => {
  localStorage.removeItem("token");
  router.push("/login");
};
</script>

<template>
  <div class="nav-container">
    <nav class="side-nav" :class="{ collapsed: isCollapsed }">
      <!-- <div class="logo" @click="handleCollapse"> -->
      <!-- 	<img src="../../assets/Onize.png" alt="Logo" /> -->
      <!-- </div> -->
      <div class="nav-items">
        <div class="logo" @click="handleCollapse">
          <img src="../../assets/Onized.png" alt="Logo" />
          <img class="half" src="../../assets/half-circle.png" alt="Logo" />
        </div>
        <router-link to="/" style="text-decoration: none">
          <i class="pi pi-home"></i>
          <span>Overview</span>
        </router-link>
        <router-link to="/Task" style="text-decoration: none">
          <i class="pi pi-check-circle"></i>
          <span>Tasks</span>
        </router-link>
        <router-link to="/Dashboard" style="text-decoration: none">
          <!-- <i class="pi pi-user"></i> -->
          <i class="pi pi-spin pi-cog"></i>
          <span>Account</span>
        </router-link>
        <div class="logout" @click="logout">
          <i class="pi pi-sign-out"></i>
          <span v-show="!isCollapsed">Logout</span>
        </div>
      </div>
    </nav>
  </div>
</template>

<style scoped>
.half {
  position: fixed;
  transition: transform 0.4s ease;
}

.half:hover {
  position: fixed;
  transform: rotate(370deg) scale(1.4);
}

.nav-container {
  display: flex;
}

.side-nav {
  position: fixed;
  width: 247px;
  height: 100vh;
  background-color: #0b0a0b;
  color: #fff;
  overflow-x: hidden;
  transition: width 0.5s ease-in-out;
  padding-top: 20px;
  display: flex;
  flex-direction: column;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.2);
  border-right: 1px solid #222;
}

.side-nav.collapsed {
  width: 80px;
}

.logo {
  /* text-align: center; */
  margin-bottom: 17px;
  cursor: pointer;
  /* background-color: white; */
  max-width: fit-content;
  justify-self: center;
  align-self: left;
  border-radius: 7px;
  /* margin-left: 22px; */
}

.logo img {
  width: 30px;
  margin-right: 0;
  padding: 0;
}

.nav-items {
  display: flex;
  flex-direction: column;
  margin-left: 20px;
  height: 100%;
  gap: 10px;
}

.nav-items a {
  padding: 10px;
  text-decoration: none;
  color: #fff;
  display: flex;
  align-items: center;
  transition: all 0.3s ease-in-out;
}

.nav-items i {
  margin-right: 10px;
  font-size: 1.5rem;
}

.side-nav span {
  transition: opacity 0.5s ease-in-out;
}

.nav-items div {
  padding: 10px;
  display: flex;
  align-items: center;
  transition: all 0.3s ease-in-out;
}

.nav-items div i {
  margin-right: 10px;
  font-size: 1.5rem;
}

div span {
  transition: opacity 0.5s ease-in-out;
}

.collapsed span {
  opacity: 0;
}

.collapsed .nav-items a,
.collapsed .nav-items div {
  padding: 10px 5px;
}

.collapsed .nav-items span,
.collapsed .nav-items div span {
  opacity: 0;
}

.logout {
  margin-top: auto;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
  cursor: pointer;
  color: #fff;
  transition: all 0.3s ease-in-out;
}
</style>

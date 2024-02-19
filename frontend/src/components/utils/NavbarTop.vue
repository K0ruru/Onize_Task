<script setup>

// primevue
import "primeicons/primeicons.css";
import Dropdown from "primevue/dropdown";
import { ref, onMounted } from "vue";

const filter = ref("");

// ngambil date
const currentDate = ref(new Date());
onMounted(() => {
  // Update the current date every second
  setInterval(() => {
    currentDate.value = new Date();
  }, 1000);
});

// Define the options for the dropdown
const filterOptions = [{ name: "Week" }, { name: "Day" }, { name: "Month" }];
</script>

<template>
  <!-- navbar-top -->
  <div class="nav-top">
    <div class="left-nav">
      <!-- <p>Jan 2024</p> -->
      <!-- <p>{{ currentDate.toLocaleString() }}</p> -->
      <p>
        {{
          currentDate.toLocaleString("en-US", {
            month: "short",
            day: "2-digit",
            year: "numeric",
          })
        }}
      </p>
      <Dropdown v-model="filter" :options="filterOptions" optionLabel="name" placeholder="Week"
        class="w-full md:w-14rem" />

      <div class="filter"><i class="pi pi-tag"></i>Filter</div>
    </div>
    <div class="right-nav">
      <i class="pi pi-chevron-left"></i>
      <i class="pi pi-calendar"></i>
      <i class="pi pi-chevron-right"></i>
    </div>
  </div>
</template>

<style scoped>
.filter {
  display: flex;
  align-items: center;
  gap: 7px;
  cursor: pointer;
}

.nav-top {
  display: flex;
  justify-content: space-between;
  width: 100%;
  border-bottom: 1px solid #222;
  padding: 1.7%;
  align-items: center;
  height: 60px;
}

.left-nav {
  display: flex;
  align-items: center;
  gap: 10px;
}

.right-nav {
  display: flex;
  gap: 10px;
}
.right-nav i {
  border-radius: 7px;
  font-size: 12px;
  background-color: rgba(90, 90, 90, 0.2);
  padding: 9px;
}
</style>

<script setup>
import "primeicons/primeicons.css";
import Dropdown from 'primevue/dropdown';
import { ref, onMounted } from "vue";



// ngambil date
const currentDate = ref(new Date());
onMounted(() => {
  // Update the current date every second
  setInterval(() => {
    currentDate.value = new Date();
  }, 1000);
});

// Define the options for the dropdown
const cities = [
  { name: 'Week' },
  { name: 'Day' },
  { name: 'Month' }
];

</script>

<template>
  <!-- navbar-top -->
  <div class="nav-top">
    <div class="left-nav">
      <!-- <p>Jan 2024</p> -->
      <!-- <p>{{ currentDate.toLocaleString() }}</p> -->
      <p>{{ currentDate.toLocaleString('en-US', { month: 'short', day: '2-digit', year: 'numeric' }) }}</p>
      <Dropdown v-model="selectedCity" :options="cities" optionLabel="name" placeholder="Week"
        class="w-full md:w-14rem" />
      <i class="pi pi-tag"></i>Filter
    </div>
    <div class="right-nav">
      <i class="pi pi-circle-fill"></i>
    </div>
  </div>
</template>

<style scoped>
.nav-top {
  display: flex;
  justify-content: space-between;
  width: 100%;
  border-bottom: 1px solid #222;
  padding: 1.7%;
}

.left-nav {
  display: flex;
  align-items: center;
  gap: 10px;
}

.right-nav {
  color: #6427aa;
}
</style>


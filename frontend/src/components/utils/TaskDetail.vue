<script setup>
import { defineEmits } from "vue";
const emits = defineEmits(["closeDetailModal"]);

const closeDetailModal = () => {
  emits("closeDetailModal");
};


</script>

<template>
  <div class="add-form-overlay">
    <div class="add-form-container">

      <!-- header ini aa -->
      <div class="form-header">
        <h1>Meeting Bersama Asep</h1>
        <Button icon="pi pi-times" @click="closeDetailModal" class="p-button-rounded p-button-text"></Button>
      </div>

      <div class="form-container">
        <div class="start-due">
          <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis eum amet cumque dolorum, aut ipsum
            accusamus natus laudantium quisquam deleniti placeat consequuntur tenetur debitis quo officia nam fugit ullam,
            quibusdam necessitatibus voluptate quaerat, quasi maiores aspernatur possimus. Reprehenderit laborum
            consequatur necessitatibus expedita fugit error tempore deleniti voluptates.</p>
        </div>
      </div>

      <div class="button-form">
        <button class="button-cencel" @click="closeDetailModal">Cencel</button>
        <button class="button-save">Done</button>
      </div>

    </div>
  </div>
</template>

<style scoped>
.button-form {
  display: flex;
  justify-content: flex-end;
  width: 100%;
  padding: 1.5rem;
  gap: 1rem;
}

.button-form .button-save {
  width: 20%;
  background: #6427aa;
  border: none;
  padding: 10px 5px;
  border-radius: 5px;
  transition: all 0.2s;
}

.button-form .button-save:hover {
  background: #6430ab;
}

.button-form .button-cencel {
  width: 20%;
  background: #0a0b0a;
  border: 1px solid #6427aa;
  padding: 10px 5px;
  border-radius: 5px;
  transition: all 0.2s;
}

.button-form .button-cencel:hover {
  border: 1px solid #ff0000;
}

.form-container {
  width: 100%;
  padding: 0rem 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.form-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-shrink: 0;
  width: 100%;
  height: 70px;
  padding: 1.5rem;
}

.form-header i {
  cursor: pointer;
}


.form-header h1 {
  font-size: 1.4rem;
}

.add-form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  /* background: rgba(0, 0, 0, 0.5); */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}

.add-form-container {
  /* background: #0b0a0b; */
  /* padding: 10px; */
  background-color: rgba(0, 0, 0, 0.7);
  border-radius: 10px;
  padding: 20px;
  /* backdrop-filter: blur(10px); */
  border-radius: 8px;
  border: 1px solid #222;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>




<script setup>
import InputText from 'primevue/inputtext';
import Editor from 'primevue/editor';
import Calendar from 'primevue/calendar';
import { defineEmits } from "vue";
const emits = defineEmits(["closeFormModal"]);

const closeFormModal = () => {
  emits("closeFormModal");
};


</script>

<template>
  <div class="add-form-overlay">
    <div class="add-form-container">

      <!-- header ini aa -->
      <div class="form-header">
        <h1>Create Task</h1>
        <!-- <i class="pi pi-times" @click="closeFormModal"></i> -->
        <Button icon="pi pi-times" @click="closeFormModal" class="p-button-rounded p-button-text"></Button>
        <!-- <Button icon="pi pi-times" @click="emits('closeFormModal')" class="p-button-rounded p-button-text"></Button> -->
      </div>

      <!-- form ini aa -->
      <div class="form-container">
        <!-- nama div -->
        <div class="name-input">
          <label for="taskname">Task Name</label>
          <InputText id="taskname" type="text" v-model="value" />
        </div>

        <!-- deskripsi div -->
        <div class="description-input">
          <label for="description">Description</label>
          <Editor id="description" v-model="value" editorStyle="height: 320px" />
        </div>

        <div class="start-due">
          <div class="start">
            <label for="start">Start</label>
            <!-- <input type="date"> -->
            <Calendar v-model="date" />
          </div>

          <div class="due">
            <label for="due">Due</label>
            <!-- <input type="date"> -->
            <Calendar v-model="date" />
          </div>
        </div>
      </div>

      <div class="button-form">
        <button class="button-cencel" @click="closeFormModal">Cencel</button>
        <button class="button-save">Save</button>
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

.start-due {
  display: flex;
  width: 100%;
  justify-content: flex-start;
  align-items: center;
  gap: 1.5rem;
  padding: 1.5rem 0rem;
}

.start {
  width: 50%;
  display: flex;
  flex-direction: column;
  gap: 7px;
}

.due {
  width: 50%;
  display: flex;
  flex-direction: column;
  gap: 7px;
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

.description-input {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 7px;
  margin-top: 10px;
}

.name-input {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 7px;
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
  background: #0b0a0b;
  /* padding: 10px; */
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}


/* prime-vue-custom */
label {
  font-size: 0.8rem;
  font-weight: 500;
  color: #fff;
}

.p-inputtext:enabled {
  border-color: #52525b;
  width: 100%;
}

.p-inputtext:enabled:hover {
  border-color: #52525b;
  width: 100%;
}
</style>



<script setup>
const emits = defineEmits(["closeAddForm"]);

</script>

<template>
  <div class="add-form-overlay">
    <div class="add-form-container">

      <!-- header ini aa -->
      <div class="form-header">
        <h1>Create Task</h1>
        <p>X</p>
      </div>

      <!-- form ini aa -->
      <div class="form-container">
        <!-- nama div -->
        <div class="name-input">
          <p>Task Name</p>
          <input type="text">
        </div>

        <!-- deskripsi div -->
        <div class="description-input">
          <p>Description</p>
          <textarea></textarea>
        </div>

        <div class="start-due">
          <div class="start">
            <p>Start Date</p>
            <input type="date">
          </div>

          <div class="due">
            <p>Due Date</p>
            <input type="date">
          </div>
        </div>
      </div>

      <div class="button-form">
        <button>Cencel</button>
        <button>Save</button>
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

.button-form button {
  width: 20%;
  background: #6427aa;
  border: none;
  padding: 4px 5px;
  border-radius: 5px;
}

.start-due {
  display: flex;
  width: 100%;
  justify-content: space-between;
}

.form-container {
  width: 100%;
  padding: 1.5rem;
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

.name-input {
  width: 100%;
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
  background: rgba(0, 0, 0, 0.5);
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
</style>


<script setup>

import axios from "axios";
import {ref} from "vue";

const props = defineProps({
  email: String,
  password: String,
})

const statusText = ref('');

const handleSubmit = () => {
  axios.post("/api/v1/login", {
    email: props.email,
    password: props.password
  }).then(res => {
    statusText.value = "Login successful!";
  }).catch(err => {
    console.log(err)
  })

};

</script>

<template>

  <h1>{{statusText}}</h1>

  <form @submit.prevent="handleSubmit">
    <div class="form-group">
      Email: <input type="email" class="form-control" id="email" v-model="email">
    </div>
    <div class="form-group">
      Password: <input type="password" class="form-control" id="password" v-model="password">
    </div>
    <br/>
    <br/>
    <br/>
    <button type="submit">Submit</button>
  </form>

</template>

<style scoped>
.form-group {
  margin-bottom: 1rem;
}
</style>

<script setup>
const navigateToWords = () => {
  window.location.hash = '#/words';
};
</script>


<template>
  <div class="title">
    <h1>Welcome to the Word Remember Website</h1>
  </div>
  <button v-on:click="navigateToWords">Get Started</button>
</template>


<style>
.title {
  text-align: center;
  color: #7393a7;
}
button {
  display: block;
  margin: 20px auto;
  padding: 10px 20px;
  font-size: 16px;
  background-color: #7393a7;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
button:hover {
  background-color: #5a6b85;
}
</style>

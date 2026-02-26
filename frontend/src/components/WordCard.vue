<script setup>

const props = defineProps(["word","discription"])

</script>

<template>
    <div class="rectangle">
        <b>{{ props.word }}</b> 
        <p v-for="d in props.discription" :key="d">{{ d }}</p>
    </div>
</template>

<style>
    .rectangle{
        width: 150px;
        background-color: #7393a7;
        color:#d0e7ef;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        border-radius: 5px;
        padding: 5px;
        margin: 10px;
    }
    .rectangle b{
        margin-bottom: 1px;
    }
    .rectangle p{
        margin: 0;
        font-size: 14px;
    }
</style>
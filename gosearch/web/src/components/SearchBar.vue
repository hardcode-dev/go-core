<template>
  <div>
    <h2>GoSearch</h2>
    <input class="search-input" type="text" v-model="query" @keypress.enter="search" />
    <button @click="search">Найти</button>
  </div>
</template>

<script>
export default {
  name: "SearchBar",
  data() {
    return {
      query: ""
    }
  },
  methods: {
    search() {
      this.$emit("search", this.query)
    }
  }
}
</script>

<style scoped lang="scss">
.search-input {
  width: 500px;
  margin-right: 10px;
}
</style>

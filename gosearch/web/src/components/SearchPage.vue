<template>
  <div>
    <div class="search-bar" :class="{ 'search-top': results && results.length > 0 }">
      <search-bar @search="search"></search-bar>
    </div>
    <search-results class="search-results" v-if="results" :results="results"></search-results>
  </div>
</template>

<script>
import SearchBar from "@/components/SearchBar"
import SearchResults from "@/components/SearchResults"

export default {
  name: "SearchPage",
  components: {
    SearchBar,
    SearchResults
  },
  data() {
    return {
      results: []
    }
  },
  methods: {
    search(query) {
      let url = "http://" + window.location.hostname + "/search/" + query
      fetch(url)
        .then(response => response.json())
        .then(data => (this.results = data))
    }
  }
}
</script>

<style scoped lang="scss">
.search-bar {
  height: 70vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.search-top {
  height: 10vh;
}

.search-results {
  display: flex;
  align-items: flex-start;
  justify-content: center;
}
</style>

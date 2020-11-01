<template>
  <div>
    <div>
      <p v-for="(result, index) in results" :key="result.ID">
        {{ index }}. <a :href="result.URL" target="_blank">{{ result.URL }}</a> <br /><b>{{ result.Title }}</b>
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: "SearchResults",
  props: {
    results: {
      type: Array,
      default: () => []
    }
  }
}
</script>

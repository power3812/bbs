<template>
  <div>
    <form v-on:submit="postBody" class="flex justify-center">
      <textarea v-model="body"><input type="text"/></textarea>
      <button>実行</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'CommonTextForm',
  data() {
    return {
      message: null,
      body: null
    }
  },
  methods: {
    async postBody() {
      await this.$axios.$post(`/api/v1/posts`, {
        body: this.body
      }).then((res) => {
        this.message = res.message
        this.body = null
      })
    }
  }
}
</script>

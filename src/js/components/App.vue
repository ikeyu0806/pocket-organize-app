<template>
  <section>
    <b-button type="is-primary" @click="clickMe">Click Me</b-button>
  </section>
</template>

<script>
export default {
  name: "App",
  methods: {
    clickMe() {
        this.$buefy.notification.open('Clicked!!')
    }
  }
}
</script>

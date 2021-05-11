<template>
  <section>
    <b-field label="Pocketのコンシューマキーを入力してください">
      <b-input v-model="consumerKey"></b-input>
      <b-button type="is-info" @click="send">送信</b-button>
    </b-field>
  </section>
</template>

<script>
export default {
  consumerKey: "",
  methods: {
    send() {
        this.$buefy.notification.open('送信しました')
    }
  }
}
</script>

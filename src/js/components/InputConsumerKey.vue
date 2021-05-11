<template>
  <section class="section">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-half">
          <b-field label="Pocketのコンシューマキーを入力してください">
            <b-input v-model="consumerKey"></b-input>
            <b-button type="is-info" @click="send">送信</b-button>
          </b-field>
        </div>
      </div>
    </div>
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

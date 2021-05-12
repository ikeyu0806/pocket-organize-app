<template>
  <section class="section">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-half">
          <b-field label="Pocketのコンシューマキーを入力してください">
            <b-button type="is-info" @click="send">Pocket認証ページを開く</b-button>
          </b-field>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  methods: {
    send() {
      this.$buefy.notification.open('送信しました');
    }
  }
}
</script>

<template>
  <section class="section">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-half">
          <b-field label="">
            <b-button type="is-info" @click="open_pocket_auth_page">Pocket認証ページを開く</b-button>
          </b-field>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
import appEnv from 'appEnv'

export default {
  methods: {
    open_pocket_auth_page() {
      axios.get(`${appEnv.apiUrl}/get_request_token`)
     .then(response => {
      console.log('status:', response.status)
      console.log('body:', response.data)
      localStorage.setItem('pocket_request_token', response.data.code)
      window.open(`https://getpocket.com/auth/authorize?request_token=${response.data.code}&redirect_uri=${appEnv.apiUrl}`, '_blank')
     }).catch(err => {
      console.log('err:', err);
     })
    }
  }
}
</script>

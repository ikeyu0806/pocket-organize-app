<template>
  <section class="section">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-half">
          <b-field label="">
            <b-button type="is-info" @click="get_access_token">認証を完了する</b-button>
          </b-field>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  methods: {
  },
  created (){
    console.log("mounted")
    const code = localStorage.getItem('pocket_request_token')
    axios.get(`${process.env.HOST_URL}/get_access_token?code=${code}`)
    .then(response => {
      console.log('status:', response.status)
      console.log('body:', response.data)
      localStorage.setItem('pocket_access_token', response.data.access_token)
      localStorage.removeItem('pocket_request_token')
    }).catch(err => {
      console.log('err:', err);
    })
  }
}
</script>

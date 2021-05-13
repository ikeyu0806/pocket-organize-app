<template>
  <div>
    <div v-if='authStatus == "requestTokenRequired"'><start-auth></start-auth></div>
    <div v-else-if='authStatus == "accessTokenRequired"'><complete-auth></complete-auth></div>
    <div v-else-if='authStatus == "authComplete"'><organize-pocket></organize-pocket></div>
  </div>
</template>

<script>
import StartAuth from './StartAuth'
import CompleteAuth from './CompleteAuth'
import OrganizePocket from './OrganizePocket'

export default {
  components: { StartAuth, CompleteAuth, OrganizePocket },
  data: function() {
          return {
            authStatus: ""
      }
  },
  mounted() {
    if(localStorage.getItem('pocket_request_token') === null && localStorage.getItem('pocket_access_token') === null) {
      this.authStatus = "requestTokenRequired"
    } else if (localStorage.getItem('pocket_access_token') != null) {
      this.authStatus = "authComplete"
    } else if (localStorage.getItem('pocket_request_token') != null) {
      this.authStatus = "accessTokenRequired"
    }
  }
}

</script>

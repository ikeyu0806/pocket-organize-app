import Vue from 'vue'
import VueRouter from 'vue-router'
import routes from './routes.js'
import App from './components/App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

Vue.use(Buefy)

const router = new VueRouter({
  routes
})

new Vue({
  el: '#app',
  router: router,
  components: {
    App
  },
  render(h) {
    return h('app')
  }
})

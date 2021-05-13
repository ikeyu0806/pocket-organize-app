import Vue from 'vue'
import App from './components/App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'

Vue.use(Buefy)

new Vue({
  el: '#app',
  components: {
    App
  },
  render(h) {
    return h('app')
  }
})

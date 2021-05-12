<template>
  <div>
    <pocket-auth></pocket-auth>
    <complete-auth></complete-auth>
  </div>
</template>

<script>
import PocketAuth from './PocketAuth'
import CompleteAuth from './CompleteAuth'

export default {
  components: { PocketAuth, CompleteAuth },
  name: "App"
}

</script>

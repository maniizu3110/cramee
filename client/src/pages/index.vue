<template>
  <div>
    <v-sheet>
      <v-container class="py-6 pt-lg-15 text-center">
        <h1 class="text-h4 text-sm-h3 text-md-h2 text-lg-h1">
          Let's find <span class="primary--text">New Style</span>
        </h1>
        <h2 class="text-h6 text-sm-h5 mt-4 w-full w-md-8-12 w-xl-half mx-auto">
          もっと柔軟に、もっと効率よく。<br />自分らしい教え方・学び方を見つけよう
        </h2>
        <div class="mt-8">
          <v-btn x-large class="my-1 mx-sm-1 w-full w-sm-auto" color="primary" to="/teacher/auth/signin"
            >講師ログイン</v-btn
          >
          <v-btn x-large class="my-1 mx-sm-1 w-full w-sm-auto" to="/student/auth/signin"
            >生徒ログイン</v-btn
          >
        </div>
      </v-container>
    </v-sheet>

    <Partners />
  </div>
</template>

<script>
import Partners from "@/components/landing/Partners.vue";

export default {
  layout: "landing",
  components: {
    Partners,
  },
};
</script>

<template>
  <div>
    <v-snackbar v-model="snackbar" :timeout="timeout">
      {{ errorMessage }}
    </v-snackbar>
    <v-card class="text-center pa-1">
      <v-card-title class="justify-center display-1 mb-2">{{
        $t("register.title")
      }}</v-card-title>
      <v-card-subtitle>{{ $t("register.detail") }}</v-card-subtitle>

      <v-card-text>
        <v-form ref="form" v-model="isFormValid" lazy-validation>
          <v-text-field
            v-model="phone"
            hint="ハイフン(-)なしで入力してください"
            :rules="[rules.required]"
            :validate-on-blur="false"
            :error="errorPhone"
            :error-messages="errorPhoneMessage"
            :label="$t('register.phone')"
            name="name"
            outlined
            @keyup.enter="submit"
            @change="resetErrors"
          ></v-text-field>

          <v-text-field
            v-model="email"
            :rules="[rules.required]"
            :validate-on-blur="false"
            :error="errorEmail"
            :error-messages="errorEmailMessage"
            :label="$t('register.email')"
            name="email"
            outlined
            @keyup.enter="submit"
            @change="resetErrors"
          ></v-text-field>

          <v-text-field
            v-model="password"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            :rules="[rules.required]"
            :type="showPassword ? 'text' : 'password'"
            :error="errorPassword"
            :error-messages="errorPasswordMessage"
            :label="$t('register.password')"
            name="password"
            outlined
            @change="resetErrors"
            @keyup.enter="submit"
            @click:append="showPassword = !showPassword"
          ></v-text-field>

          <v-btn
            :loading="isLoading"
            :disabled="isSignUpDisabled"
            block
            x-large
            color="primary"
            @click="submit"
            >{{ $t("register.button") }}</v-btn
          >
        </v-form>
      </v-card-text>
    </v-card>

    <div class="text-center mt-6">
      {{ $t("register.account") }}
      <router-link
        :to="localePath('/teacher/auth/signin')"
        class="font-weight-bold"
      >
        {{ $t("register.signin") }}
      </router-link>
    </div>
    <div class="text-center mt-6">
      {{ $t("register.teacher") }}
      <router-link
        :to="localePath('/teacher/auth/signup')"
        class="font-weight-bold"
      >
        {{ $t("register.to_teacher_signup") }}
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  layout: "auth",
  data() {
    return {
      //snackbar
      snackbar: false,
      timeout: 2000,
      errorMessage: "",
      // sign up buttons
      isLoading: false,
      isSignUpDisabled: false,

      // form
      isFormValid: true,
      email: "",
      password: "",
      phone: "",

      // form error
      errorPhone: false,
      errorEmail: false,
      errorPassword: false,
      errorPhoneMessage: "",
      errorEmailMessage: "",
      errorPasswordMessage: "",
      // show password field
      showPassword: false,
      // input rules
      rules: {
        required: (value) => (value && Boolean(value)) || "Required",
      },
    };
  },
  methods: {
    submit() {
      if (this.$refs.form.validate()) {
        this.isLoading = true;
        this.isSignUpDisabled = true;
        this.signUp(this.phone, this.email, this.password);
      }
    },
    signUp(phone, email, password) {
      this.$axios
        .post("/v1/sign-teacher/with-zoom", {
          phone_number: phone,
          email: email,
          hashed_password: password,
        })
        .then((res) => {
          this.isLoading = false;
          this.isSignUpDisabled = false;
          this.$router.push("/teacher/auth/signin")
        })
        .catch((e) => {
          this.errorMessage = e.message;
          this.snackbar = true;
          this.isLoading = false;
          this.isSignUpDisabled = false;
        });
    },
    resetErrors() {
      this.errorPhone = false;
      this.errorEmail = false;
      this.errorPassword = false;
      this.errorPhoneMessage = "";
      this.errorEmailMessage = "";
      this.errorPasswordMessage = "";
    },
  },
};
</script>

export const state = () => ({
  isLogin: false,
  isStudent: false,
  id: null,
});

export const getters = {
  getId: (state) => state.id,
};

export const mutations = {
  loginStudent(state, payload) {
    state.isLogin = true;
    state.isStudent = true;
    state.id = payload.id;
  },
  loginTeacher(state, payload) {
    state.isLogin = true;
    state.isStudent = false;
    state.id = payload.id;
  },
  logout(state) {
    state.isLogin = false;
    state.isStudent = false;
    state.id = null;
  },
};

export const actions = {
  async loginStudent(context, payload) {
    await this.$axios
      .post("/v1/sign-student/login", {
        email: payload.email,
        password: payload.password,
      })
      .then((res) => {
        context.commit("loginStudent", res.data.student);
        localStorage.setItem("Bearer", res.data.access_token);
      });
  },
  async loginTeacher(context, payload) {
    let data;
    await this.$axios
      .post("/v1/sign-teacher/login", {
        email: payload.email,
        password: payload.password,
      })
      .then((res) => {
        context.commit("loginTeacher", res.data.teacher);
        localStorage.setItem("Bearer", res.data.access_token);
        data = res.data;
      });
    return data;
  },
  async logout(context) {
    context.commit("logout");
    localStorage.removeItem("Bearer");
    return true;
  },
};

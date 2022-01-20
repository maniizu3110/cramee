import menuPages from "./menus/pages.menu";

export default {
  menues: {
    student: [
      {
        text: "Home",
        key: "",
        items: [
          {
            icon: "mdi-home",
            key: "menu.home",
            text: "TOP",
            link: "/",
          },
        ],
      },
      {
        text: "Auth",
        key: "",
        items: [
          {
            icon: "mdi-account-plus",
            key: "menu.signup",
            text: "signup",
            link: "/student/auth/signup",
          },
          {
            icon: "mdi-login",
            key: "menu.signin",
            text: "signin",
            link: "/student/auth/signin",
          },
          {
            icon: "mdi-logout",
            key: "menu.logout",
            text: "logout",
            link: "/student/auth/logout",
          },
        ],
      },
      {
        text: "Schedule",
        key: "",
        items: [
          {
            icon: "mdi-calendar",
            key: "menu.schedule",
            text: "schedule",
            link: "/student/auth/signup",
          },
        ],
      },
    ],
    teacher: [
      {
        text: "Home",
        key: "",
        items: [
          {
            icon: "mdi-view-dashboard-outline",
            key: "menu.dashboard",
            text: "Dashboard",
            link: "/",
          },
          {
            icon: "mdi-file-outline",
            key: "menu.blank",
            text: "Blank Page",
            link: "/blank",
          },
        ],
      },
      {
        text: "Pages",
        key: "menu.pages",
        items: menuPages,
      },
      {
        text: "Landing Pages",
        items: [
          {
            icon: "mdi-airplane-landing",
            key: "menu.landingPage",
            text: "Landing Page",
            link: "/landing",
          },
          {
            icon: "mdi-cash-usd-outline",
            key: "menu.pricingPage",
            text: "Pricing Page",
            link: "/landing/pricing",
          },
        ],
      },
    ],
    default: [
      {
        text: "Home",
        key: "",
        items: [
          {
            icon: "mdi-home",
            key: "menu.home",
            text: "TOP",
            link: "/",
          },
        ],
      },
      {
        text: "Auth",
        key: "",
        items: [
          {
            icon: "mdi-book",
            key: "menu.loginStudent",
            text: "login-student",
            link: "/student/auth/signin",
          },
          {
            icon: "mdi-human-male-board",
            key: "menu.loginTeacher",
            text: "login-teachr",
            link: "/teacher/auth/signin",
          },
        ],
      },
    ],
  },
};

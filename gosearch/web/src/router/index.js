import Vue from "vue"
import VueRouter from "vue-router"
import SearchPage from "../components/SearchPage"

Vue.use(VueRouter)

const routes = [
  {
    path: "/",
    name: "SearchPage",
    component: SearchPage,
    meta: {
      title: "GoSearch"
    }
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title
  next()
})

export default router

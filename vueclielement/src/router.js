import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Layout from './views/Layout.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
        path: '/layout',
        component: Layout
    },
    {
        path: '/container',
        component: () => import('./views/Container.vue')
    },
    {
        path: '/input',
        component: () => import('./views/Input.vue')
    },
    {
      path: '/select',
      component: () => import('./views/Select.vue')
    },
    {
      path: '/datepicker',
      component: () => import('./views/DatePicker.vue')
    },
    {
      path: '/form',
      component: () => import('./views/Form.vue')
    },
    {
      path: '/table',
      component: () => import('./views/Table.vue')
    },
    {
      path: '/dialog',
      component: () => import('./views/Dialog.vue')
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    }
  ]
})

// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueSidebarMenu from 'vue-sidebar-menu'
import 'vue-sidebar-menu/dist/vue-sidebar-menu.css'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { dom } from '@fortawesome/fontawesome-svg-core'
import VueNavigationBar from 'vue-navigation-bar'
import treeNav from 'vue-tree-nav'

Vue.component('vue-navigation-bar', VueNavigationBar)

dom.watch()


Vue.use(VueSidebarMenu)
Vue.config.productionTip = false
Vue.component('font-awesome-icon', FontAwesomeIcon)


/* eslint-disable no-new */
// new Vue({
//   el: '#app',
//   router,
//   components: { App },
//   template: '<App/>'
// })

new Vue({
  components: {
    'vue-tree-nav': treeNav,
    App
  },
  data: {
    left: [
      {
        label: 'Home',
        icon: 'home',
        href: '#/home'
      }, {
        label: 'Animals',
        children: [
          {
            label: 'Elephant',
            href: '#/animals/elephant'
          }, {
            label: 'Lion',
            href: '#/animals/lion'
          }, {
            label: 'Bear',
            href: '#/animals/bear'
          }, {
            label: 'Eagle',
            href: '#/animals/eagle'
          }, {
            label: 'Wolf',
            href: '#/animals/wolf'
          }

        ]
      }
    ],
    right: [
      {
        href: 'https://github.com/marcodpt/vue-tree-nav',
        icon: 'brands/github'
      }
    ]
  }
}).$mount('#app')

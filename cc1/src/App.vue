/* eslint-disable no-alert, no-console */
<template>
  <div id="app">
    <vue-navigation-bar :options="options" />
    <sidebar-menu :menu="menu" />
  </div>
</template>
<script>
import { SidebarMenu } from 'vue-sidebar-menu'
export default {
  name: 'App',
  components: {
    SidebarMenu
  },
  data() {
      return {
        menu: [
          {
            header: false,
            hiddenOnCollapse: true
          },
          {
            href: '/',
            title: 'Future Meetings',
          },
          {
            href: '/past',
            title: 'Previous Meetings',
          },
          {
            href: '/mindmap',
            title: 'Mind Map',
          }
        ],
        navbarData: {
                    elementId: 'main-navbar',
                    isUsingVueRouter: true,
                    mobileBreakpoint: 992,
                    brandImagePath: '../src/assets/logo.png',
                    brandImageAltText: 'vue-navigation-bar',
                    collapseButtonStyle: 'dark',
                    ariaLabelMainNav: 'Main Navigation',
                    menuOptionsLeft: [
                        {
                            type: 'link',
                            text: 'Why Dunder Mifflin',
                            path: '/why',
                            subMenuOptions: [
                                {
                                    type: 'link',
                                    text: 'About',
                                    path: '#',
                                },
                                {
                                    type: 'hr',
                                },
                                {
                                    type: 'link',
                                    text: 'Locations',
                                    path: '/locations',
                                },
                                {
                                    type: 'link',
                                    text: 'Blog',
                                    path: '/blog',
                                },
                            ]
                        },
                        {
                            type: 'link',
                            text: 'Contact',
                            path: '/contact',
                            subMenuOptions: [
                                {
                                    type: 'link',
                                    text: 'Customer Service',
                                    path: '/customer-service',
                                },
                                {
                                    type: 'link',
                                    text: 'Accounting',
                                    path: '/accounting',
                                },
                                {
                                    type: 'link',
                                    text: 'Reception',
                                    path: '/reception',
                                },
                            ]
                        },
                        {
                            type: 'link',
                            text: 'Pricing',
                            path: '/pricing',
                        },
                    ],
                    menuOptionsRight: [
                        {
                            type: 'button',
                            text: 'Signup',
                            path: '/signup',
                            class: 'button-red'
                        },
                        {
                            type: 'button',
                            text: 'Login',
                            path: '/login',
                        }
                    ]
                }
      }
  }
}
</script>
<style lang-"sass">
.navbar-item {
  color:#4a4a4a !important;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
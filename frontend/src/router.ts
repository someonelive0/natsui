import { createRouter, createWebHashHistory } from 'vue-router'
// import Bar from './components/Bar.vue'
// import Foo from './components/Foo.vue'
import Dashboard from './components/Dashboard.vue'



const routes = [
    { path: '/', name: 'Dashboard', component: Dashboard },
    // { path: '/foo', name: 'Foo', component: Foo,
    //   props: true, mete: { title: 'foo' }
    // },
    // { path: '/bar', name: 'Bar', component: Bar,
    //   props: true, mete: { title: 'bar' }
    // },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router

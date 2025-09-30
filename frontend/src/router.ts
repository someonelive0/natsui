import { createRouter, createWebHashHistory } from 'vue-router'
// import Bar from './components/Bar.vue'
// import Foo from './components/Foo.vue'
import HelloWorld from './components/HelloWorld.vue'



const routes = [
    { path: '/', name: 'HelloWorld', component: HelloWorld },
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

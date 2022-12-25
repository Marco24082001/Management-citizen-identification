// import Login from '@/pages/login/index.vue'
// // import home from '../pages/home/index.vue'

// const routes = [
//     {
//         path: '/login',
//         name: 'Login',
//         component: Login ,
//     }
// ]

// export default routes


import DashboardLayout from "@/layout/dashboard/DashboardLayout.vue";
// GeneralViews


// Admin pages
import Login from "@/pages/login/index.vue";
import identications from "@/pages/identications/index.vue";
// import UserProfile from "@/pages/UserProfile.vue";
// import Notifications from "@/pages/Notifications.vue";
// import Icons from "@/pages/Icons.vue";
// import Maps from "@/pages/Maps.vue";
// import Typography from "@/pages/Typography.vue";
// import TableList from "@/pages/TableList.vue";
// import Positions from "@/pages/positions/index.vue";
// import Voters from "@/pages/voters/index.vue";
// import Votes from "@/pages/votes/index.vue";
// import Electionconfigure from "@/pages/electionconfigure/index.vue";
// import NotFound from "@/pages/NotFoundPage.vue";

const routes = [
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/identications",  
    children: [
      {
        path: "identications",
        name: "identications",
        component: identications
      },
    ]
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
//   { path: "*", component: NotFound }
];

/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes;

import Detail from 'views/Detail';
import Create from 'views/Create';
import History from 'views/History';
import NotFound from 'views/NotFound';
import SendTransaction from 'views/SendTransaction';

const routes = [
  {
    path: '/detail',
    exact: true,
    auth: true,
    component: Detail,
  },
  {
    path: '/send-transaction',
    exact: true,
    auth: true,
    component: SendTransaction,
  },
  {
    path: '/history',
    exact: true,
    auth: true,
    component: History,
  },
  {
    path: '/create',
    exact: true,
    component: Create
  },
  {
    component: NotFound
  }
]

export default routes;

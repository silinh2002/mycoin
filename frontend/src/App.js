import React, { Component } from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";
import { createBrowserHistory } from "history";
import routes from "./routes";
import "styles/style.scss";
import Menu from "views/Menu";
import "bootstrap/dist/css/bootstrap.min.css";

const history = createBrowserHistory();

function RouteWithSubRoutes(route) {
  return (
    <Route
      path={route.path}
      exact={route.exact}
      render={(props) => {
        const { location: { pathname } = {} } = props;
        return pathname === "/" ? (
          <Redirect to="/create" />
        ) : (
          <route.component {...props} routes={route.routes} />
        );
      }}
    />
  );
}

export default class App extends Component {
  render() {
    return (
      <Router history={history}>
        <div className="layout">
          <Menu />
          <div className="layout__wrapper">
            <Switch>
              {routes.map((route, index) => (
                <RouteWithSubRoutes key={index} {...route} />
              ))}
            </Switch>
          </div>
        </div>
      </Router>
    );
  }
}

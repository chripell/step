package com.google.sps;

import com.googlecode.objectify.*;
import java.util.logging.Logger;
import com.google.appengine.api.datastore.DatastoreService;
import com.google.appengine.api.datastore.DatastoreServiceFactory;
import com.google.appengine.api.datastore.Entity;
import java.io.IOException;
import javax.servlet.*;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;


// NOTE!
// to run locally you need to:
// run the data store emulator somewhere. You can install it with:
//  gcloud components install cloud-datastore-emulator
// eventually login:
//  gcloud auth application-default login
// set the name of your project (here it is chripell-playground):
//  gcloud config set project chripell-playground
// and run the emulator:
//  gcloud beta emulators datastore start --host-port=localhost:8484
// Then you can go back to where you are working on your project:
//  export DATASTORE_EMULATOR_HOST=localhost:8484
//  export GOOGLE_CLOUD_PROJECT=chripell-playground
//  mvn package appengine:run

public class Init implements ServletContextListener {
  private static final Logger log = Logger.getLogger(Init.class.getName());

  public void contextInitialized(ServletContextEvent event) {
    ObjectifyService.init();
    ObjectifyService.register(com.google.sps.data.Tasko.class);
    log.info("Hello servlet World!");
  }
  public void contextDestroyed(ServletContextEvent event) {
    log.info("Goodbye servlet World!");
  }
}

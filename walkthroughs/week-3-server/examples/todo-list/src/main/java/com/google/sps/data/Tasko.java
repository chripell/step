package com.google.sps.data;

import com.googlecode.objectify.*;
import com.googlecode.objectify.annotation.*;

@Entity
public class Tasko {
  @Id Long id;
  String title;
  @Index long timestamp;

  public Tasko(String title) {
    this.title = title;
    this.timestamp = System.currentTimeMillis();
  }
  public Tasko() {}
}

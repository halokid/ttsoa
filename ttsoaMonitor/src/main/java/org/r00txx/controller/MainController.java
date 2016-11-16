package org.r00txx.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

/**
 * Created by r00xx<82049406@qq.com> on 2016/11/16.
 */

@Controller
public class MainController {

  @RequestMapping(value = "/", method = RequestMethod.GET)
  public String index() {
    return "index";
  }
}

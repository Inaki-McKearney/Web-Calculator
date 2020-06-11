package com.webcalc.power;

import org.springframework.http.*;
// import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.lang.Math; 
import java.util.Map ;


@RestController
@RequestMapping("/power")
public class PowerController {
    // @CrossOrigin()
    @GetMapping
    public ResponseEntity<?> power(
        @RequestParam Map<String,String> allParams,
        @RequestParam(required=false, value="x") String x,
        @RequestParam(required=false, value="y") String y) {

        if (x == null)
            return BadRequestError("x parameter not provided");
        if (y == null)
            return BadRequestError("y parameter not provided");
        
        allParams.remove("x");
        allParams.remove("y");
        if (!allParams.isEmpty()) {
            return BadRequestError("Unrecognised parameter(s) provided");
        }

        try {
            Long.parseLong(x);
        }
        catch (NumberFormatException e) {
            return BadRequestError("x parameter is not a valid integer");
        }

        try {
            Long.parseLong(y);
        }
        catch (NumberFormatException e) {
            return BadRequestError("y parameter is not a valid integer");
        }
      
        long answer = (long) Math.pow(Double.parseDouble(x), Double.parseDouble(y));

        if (answer == Long.MAX_VALUE)
            return BadRequestError("Result too large");
        
        Response res = new Response(
            false,
            x + "^" + y + " = " + answer,
            answer
            );

        return ResponseEntity.status(HttpStatus.OK).body(res);    
    }

    public ResponseEntity<?> BadRequestError(String message) {
        Error err = new Error(true, message);
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(err);
    }
}

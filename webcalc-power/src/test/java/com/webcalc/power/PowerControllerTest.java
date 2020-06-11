package com.webcalc.power;

import static org.hamcrest.Matchers.is;
import static org.hamcrest.Matchers.equalTo;

import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;


@RunWith(SpringRunner.class)
@SpringBootTest
@AutoConfigureMockMvc
public class PowerControllerTest {

    @Autowired
    private MockMvc mvc;

    @Test
    public void test_1() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=2&y=3").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error", is(false)))
                .andExpect(jsonPath("$.string", equalTo("2^3 = 8")))
                .andExpect(jsonPath("$.answer", is(8)));                
    }

    @Test
    public void test_2() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=-4&y=2").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error", is(false)))
                .andExpect(jsonPath("$.string", equalTo("-4^2 = 16")))
                .andExpect(jsonPath("$.answer", is(16)));                
    }

    @Test
    public void test_3() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=5&y=-3").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error", is(false)))
                .andExpect(jsonPath("$.string", equalTo("5^-3 = 0")))
                .andExpect(jsonPath("$.answer", is(0)));                
    }

    @Test
    public void test_4() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=0&y=5").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error", is(false)))
                .andExpect(jsonPath("$.string", equalTo("0^5 = 0")))
                .andExpect(jsonPath("$.answer", is(0)));                
    }

    @Test
    public void test_5() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=19&y=0").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.error", is(false)))
                .andExpect(jsonPath("$.string", equalTo("19^0 = 1")))
                .andExpect(jsonPath("$.answer", is(1)));                
    }

    @Test
    public void test_error_1() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?&y=15").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error", is(true)))
                .andExpect(jsonPath("$.message", equalTo("x parameter not provided")));
    }

    @Test
    public void test_error_2() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=i&y=15").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error", is(true)))
                .andExpect(jsonPath("$.message", equalTo("x parameter is not a valid integer")));
    }

    @Test
    public void test_error_3() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=3").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error", is(true)))
                .andExpect(jsonPath("$.message", equalTo("y parameter not provided")));
    }

    @Test
    public void test_error_4() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=45&y=s").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error", is(true)))
                .andExpect(jsonPath("$.message", equalTo("y parameter is not a valid integer")));
    }

    @Test
    public void test_error_5() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=45&y=3&z=2").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error", is(true)))
                .andExpect(jsonPath("$.message", equalTo("Unrecognised parameter(s) provided")));
    }

    @Test
    public void test_error_6() throws Exception {
        mvc.perform(MockMvcRequestBuilders.get("/power?x=45&y=33").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error", is(true)))
                .andExpect(jsonPath("$.message", equalTo("Result too large")));
    }
}


/*
mockMvc.perform(post("/books")
.content(bookInJson)
.header(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON))
.andDo(print())
.andExpect(status().isBadRequest())
.andExpect(jsonPath("$.timestamp", is(notNullValue())))
.andExpect(jsonPath("$.status", is(400)))
.andExpect(jsonPath("$.errors").isArray())
.andExpect(jsonPath("$.errors", hasSize(3)))
.andExpect(jsonPath("$.errors", hasItem("Author is not allowed.")))
.andExpect(jsonPath("$.errors", hasItem("Please provide a author")))
.andExpect(jsonPath("$.errors", hasItem("Please provide a price")));*/
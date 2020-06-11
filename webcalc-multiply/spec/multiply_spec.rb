require 'spec_helper.rb'

describe MultiplyServer do
  it "should multiply positive numbers" do
    get '/multiply?x=3&y=2'
    expect(last_response).to be_ok
    expect(last_response.body).to match %r{"error":false}
    expect(last_response.body).to match %r{"string":"3\*2=6"}
    expect(last_response.body).to match %r{"answer":6}
  end

  it "should multiply positive and negative numbers" do
    get '/multiply?x=10&y=-2'
    expect(last_response).to be_ok
    expect(last_response.body).to match %r{"error":false}
    expect(last_response.body).to match %r{"string":"10\*-2=-20"}
    expect(last_response.body).to match %r{"answer":-20}
  end

  it "should multiply negative and positive numbers" do
    get '/multiply?x=-5&y=3'
    expect(last_response).to be_ok
    expect(last_response.body).to match %r{"error":false}
    expect(last_response.body).to match %r{"string":"-5\*3=-15"}
    expect(last_response.body).to match %r{"answer":-15}
  end

  it "should multiply negative numbers" do
    get '/multiply?x=-7&y=-3'
    expect(last_response).to be_ok
    expect(last_response.body).to match %r{"error":false}
    expect(last_response.body).to match %r{"string":"-7\*-3=21"}
    expect(last_response.body).to match %r{"answer":21}
  end

  it "should multiply negative numbers" do
    get '/multiply?x=-7&y=-3'
    expect(last_response).to be_ok
    expect(last_response.body).to match %r{"error":false}
    expect(last_response.body).to match %r{"string":"-7\*-3=21"}
    expect(last_response.body).to match %r{"answer":21}
  end

  it "should require an x parameter" do
    get '/multiply?y=-3'
    expect(last_response.status).to eq 400
    expect(last_response.body).to match %r{"error":true}
    expect(last_response.body).to match %r{"message":"x parameter not provided"}
  end

  it "should require an intger for x parameter" do
    get '/multiply?x="m"&y=-3'
    expect(last_response.status).to eq 400
    expect(last_response.body).to match %r{"error":true}
    expect(last_response.body).to match %r{"message":"x parameter is not a valid integer"}
  end

  it "should require an y parameter" do
    get '/multiply?x=2'
    expect(last_response.status).to eq 400
    expect(last_response.body).to match %r{"error":true}
    expect(last_response.body).to match %r{"message":"y parameter not provided"}
  end

  it "should require an intger for y parameter" do
    get '/multiply?x=8&y="k"'
    expect(last_response.status).to eq 400
    expect(last_response.body).to match %r{"error":true}
    expect(last_response.body).to match %r{"message":"y parameter is not a valid integer"}
  end

  it "should not accept extra parameters" do
    get '/multiply?x=3&y=2&z=2'
    expect(last_response.status).to eq 400
    expect(last_response.body).to match %r{"error":true}
    expect(last_response.body).to match %r{"message":"Unexpected parameter\(s\) provided"}
  end
end
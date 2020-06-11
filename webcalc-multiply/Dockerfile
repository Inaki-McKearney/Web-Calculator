FROM ruby:2.3-alpine3.8

#RUN apt-get update -qq && apt-get install -y build-essential

ENV APP_HOME /app
WORKDIR $APP_HOME

ADD Gemfile* $APP_HOME/
RUN bundle install --without development test

ADD . $APP_HOME

EXPOSE 80

CMD ["bundle", "exec", "rackup", "-o", "0.0.0.0", "-p", "80"]
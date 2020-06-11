let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../server');
let should = chai.should();


chai.use(chaiHttp);

describe('/GET subtract', () => {
    it('it should subtract y from x', (done) => {
        chai.request(server)
            .get('/subtract?x=10&y=2')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(false);
                    res.body.should.have.property('string').eql('10-2=8');
                    res.body.should.have.property('answer').eql(8);
            done();
        });
    }); 

    it('it should subtract y from x', (done) => {
        chai.request(server)
            .get('/subtract?x=-10&y=3')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(false);
                    res.body.should.have.property('string').eql('-10-3=-13');
                    res.body.should.have.property('answer').eql(-13);
            done();
        });
    }); 

    it('it should subtract y from x', (done) => {
        chai.request(server)
            .get('/subtract?x=50&y=-10')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(false);
                    res.body.should.have.property('string').eql('50--10=60');
                    res.body.should.have.property('answer').eql(60);
            done();
        });
    }); 

    it('it should require a valid x', (done) => {
        chai.request(server)
            .get('/subtract?y=-10')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(true);
                    res.body.should.have.property('string').eql('x parameter invalid/missing');
                    res.body.should.have.property('answer').eql(0);
            done();
        });
    });

    it('it should require a valid x', (done) => {
        chai.request(server)
            .get('/subtract?x=k&y=-10')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(true);
                    res.body.should.have.property('string').eql('x parameter invalid/missing');
                    res.body.should.have.property('answer').eql(0);
            done();
        });
    });

    it('it should require a valid y', (done) => {
        chai.request(server)
            .get('/subtract?x=10')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(true);
                    res.body.should.have.property('string').eql('y parameter invalid/missing');
                    res.body.should.have.property('answer').eql(0);
            done();
        });
    });

    it('it should require a valid x', (done) => {
        chai.request(server)
            .get('/subtract?x=8&y=d')
            .end((err, res) => {
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('error').eql(true);
                    res.body.should.have.property('string').eql('y parameter invalid/missing');
                    res.body.should.have.property('answer').eql(0);
            done();
        });
    });
});
/**
 * @fileoverview gRPC-Web generated client stub for gotemplateService
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.gotemplateService = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.gotemplateService.GotemplateServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.gotemplateService.GotemplateServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.gotemplateService.GotemplateServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.gotemplateService.GotemplateServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.gotemplateService.ServiceCommand,
 *   !proto.gotemplateService.ServiceCommand>}
 */
const methodInfo_GotemplateService_Health = new grpc.web.AbstractClientBase.MethodInfo(
  proto.gotemplateService.ServiceCommand,
  /** @param {!proto.gotemplateService.ServiceCommand} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.gotemplateService.ServiceCommand.deserializeBinary
);


/**
 * @param {!proto.gotemplateService.ServiceCommand} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.gotemplateService.ServiceCommand)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gotemplateService.ServiceCommand>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gotemplateService.GotemplateServiceClient.prototype.health =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gotemplateService.GotemplateService/Health',
      request,
      metadata,
      methodInfo_GotemplateService_Health,
      callback);
};


/**
 * @param {!proto.gotemplateService.ServiceCommand} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gotemplateService.ServiceCommand>}
 *     The XHR Node Readable Stream
 */
proto.gotemplateService.GotemplateServicePromiseClient.prototype.health =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.health(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.gotemplateService;


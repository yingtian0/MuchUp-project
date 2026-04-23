from google.protobuf import descriptor_pb2, descriptor_pool, message_factory


def _build_messages():
    file_proto = descriptor_pb2.FileDescriptorProto()
    file_proto.name = "llm/v1/llm.proto"
    file_proto.package = "llm.v1"
    file_proto.syntax = "proto3"

    context_message = file_proto.message_type.add()
    context_message.name = "ContextMessage"
    for name, number, field_type in [
        ("message_id", 1, descriptor_pb2.FieldDescriptorProto.TYPE_STRING),
        ("room_id", 2, descriptor_pb2.FieldDescriptorProto.TYPE_STRING),
        ("user_id", 3, descriptor_pb2.FieldDescriptorProto.TYPE_STRING),
        ("role", 4, descriptor_pb2.FieldDescriptorProto.TYPE_STRING),
        ("content", 5, descriptor_pb2.FieldDescriptorProto.TYPE_STRING),
        ("created_at", 6, descriptor_pb2.FieldDescriptorProto.TYPE_INT64),
    ]:
        field = context_message.field.add()
        field.name = name
        field.number = number
        field.label = descriptor_pb2.FieldDescriptorProto.LABEL_OPTIONAL
        field.type = field_type

    generate_reply_request = file_proto.message_type.add()
    generate_reply_request.name = "GenerateReplyRequest"
    request_fields = [
        ("room_id", 1, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("session_id", 2, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("target_user_id", 3, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("system_prompt", 4, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("messages", 5, descriptor_pb2.FieldDescriptorProto.TYPE_MESSAGE, ".llm.v1.ContextMessage"),
        ("model", 6, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("temperature", 7, descriptor_pb2.FieldDescriptorProto.TYPE_FLOAT, ""),
        ("max_tokens", 8, descriptor_pb2.FieldDescriptorProto.TYPE_INT32, ""),
        ("metadata", 9, descriptor_pb2.FieldDescriptorProto.TYPE_MESSAGE, ".llm.v1.GenerateReplyRequest.MetadataEntry"),
    ]
    metadata_entry = generate_reply_request.nested_type.add()
    metadata_entry.name = "MetadataEntry"
    metadata_entry.options.map_entry = True
    key_field = metadata_entry.field.add()
    key_field.name = "key"
    key_field.number = 1
    key_field.label = descriptor_pb2.FieldDescriptorProto.LABEL_OPTIONAL
    key_field.type = descriptor_pb2.FieldDescriptorProto.TYPE_STRING
    val_field = metadata_entry.field.add()
    val_field.name = "value"
    val_field.number = 2
    val_field.label = descriptor_pb2.FieldDescriptorProto.LABEL_OPTIONAL
    val_field.type = descriptor_pb2.FieldDescriptorProto.TYPE_STRING
    for name, number, field_type, type_name in request_fields:
        field = generate_reply_request.field.add()
        field.name = name
        field.number = number
        field.label = (
            descriptor_pb2.FieldDescriptorProto.LABEL_REPEATED
            if name in {"messages", "metadata"}
            else descriptor_pb2.FieldDescriptorProto.LABEL_OPTIONAL
        )
        field.type = field_type
        if type_name:
            field.type_name = type_name

    token_usage = file_proto.message_type.add()
    token_usage.name = "TokenUsage"
    for name, number in [("prompt_tokens", 1), ("completion_tokens", 2), ("total_tokens", 3)]:
        field = token_usage.field.add()
        field.name = name
        field.number = number
        field.label = descriptor_pb2.FieldDescriptorProto.LABEL_OPTIONAL
        field.type = descriptor_pb2.FieldDescriptorProto.TYPE_INT32

    generate_reply_response = file_proto.message_type.add()
    generate_reply_response.name = "GenerateReplyResponse"
    response_fields = [
        ("reply_id", 1, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("room_id", 2, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("model", 3, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("content", 4, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("finish_reason", 5, descriptor_pb2.FieldDescriptorProto.TYPE_STRING, ""),
        ("created_at", 6, descriptor_pb2.FieldDescriptorProto.TYPE_INT64, ""),
        ("usage", 7, descriptor_pb2.FieldDescriptorProto.TYPE_MESSAGE, ".llm.v1.TokenUsage"),
    ]
    for name, number, field_type, type_name in response_fields:
        field = generate_reply_response.field.add()
        field.name = name
        field.number = number
        field.label = descriptor_pb2.FieldDescriptorProto.LABEL_OPTIONAL
        field.type = field_type
        if type_name:
            field.type_name = type_name

    service = file_proto.service.add()
    service.name = "LLMService"
    method = service.method.add()
    method.name = "GenerateReply"
    method.input_type = ".llm.v1.GenerateReplyRequest"
    method.output_type = ".llm.v1.GenerateReplyResponse"

    pool = descriptor_pool.DescriptorPool()
    pool.Add(file_proto)
    return {
        "ContextMessage": message_factory.GetMessageClass(pool.FindMessageTypeByName("llm.v1.ContextMessage")),
        "GenerateReplyRequest": message_factory.GetMessageClass(pool.FindMessageTypeByName("llm.v1.GenerateReplyRequest")),
        "TokenUsage": message_factory.GetMessageClass(pool.FindMessageTypeByName("llm.v1.TokenUsage")),
        "GenerateReplyResponse": message_factory.GetMessageClass(pool.FindMessageTypeByName("llm.v1.GenerateReplyResponse")),
    }


MESSAGES = _build_messages()
ContextMessage = MESSAGES["ContextMessage"]
GenerateReplyRequest = MESSAGES["GenerateReplyRequest"]
TokenUsage = MESSAGES["TokenUsage"]
GenerateReplyResponse = MESSAGES["GenerateReplyResponse"]

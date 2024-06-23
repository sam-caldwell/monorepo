/**
 * projects/exceptions/src/exception.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * This is the header file which a consumer can include to make available
 * ALL exceptions.
 *
 * By defining each preprocessor label for each EXCEPTION, the exception
 * class is made available to the build.
 *
 */
#ifndef EXCEPTIONS_H
#define EXCEPTIONS_H

#include "classes/base.h"
#include "classes/BadFilename.h"
#include "classes/BoundsCheckError.h"
#include "classes/ClosedChannelException.h"
#include "classes/FileIoError.h"
#include "classes/InvalidKeyString.h"
#include "classes/LockTimeoutExpired.h"
#include "classes/MissingInputs.h"
#include "classes/NotFound.h"
#include "classes/NullValue.h"
#include "classes/ReadOnlyKeyValue.h"
#include "classes/TypeError.h"
#include "classes/UnexpectedType.h"
#include "classes/UnsupportedFileFormat.h"
#include "classes/ValueAssignmentError.h"
#include "classes/ValueError.h"
#include "classes/ValueGetStringError.h"

#endif //EXCEPTIONS_H

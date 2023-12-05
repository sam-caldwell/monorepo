/**
 * SmartPointer/src/SmartPointer.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */

template<class ActualType>
class SmartPtr {
private:
    /**
     * pointer to actual data object.
     */
    ActualType *ptr;
public:
    /**
     * Class Constructor
     * @param p <ActualType>
     */
    explicit SmartPtr(ActualType *p = NULL) { ptr = p; }

    /**
     * Class Destructor
     */
    ~SmartPtr() { delete (ptr); }

    /**
     * Overloaded de-referencing (*) operator
     * @return ActualType reference
     */
    ActualType &operator*() { return *ptr; }
    /**
     * Overloaded arrow (->) operator allows
     * members of <ActualType> to be accessible
     * like a pointer (to allow support of classes,
     * structs or unions).
     * @return ActualType pointer.
     */
    ActualType *operator->() { return ptr; }
};
